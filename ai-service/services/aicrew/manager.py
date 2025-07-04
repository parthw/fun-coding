import asyncio
import os
from typing import Type

from crewai import LLM, Agent, Crew, Process, Task
from crewai.tools import BaseTool
from pydantic import BaseModel, Field

# Set up environment
os.environ["OPENAI_API_KEY"] = "NA"


class TopicClassificationInput(BaseModel):
    """Input model for topic classification tool."""

    message: str = Field(
        ...,
        description="The message to classify as clinical trial related or not.",
    )


class TopicClassificationTool(BaseTool):
    name: str = "TopicClassificationTool"
    description: str = "Classifies messages as clinical trial related or not based on specific keywords."
    args_schema: Type[BaseModel] = TopicClassificationInput

    async def _run(self, message: str) -> str:
        """
        Classify the message as clinical trial related or not based on clinical_keywords.
        """
        clinical_keywords = [
            "clinical trial",
            "protocol",
            "fda",
            "ema",
            "phase i",
            "phase ii",
            "phase iii",
            "randomized",
            "placebo",
            "endpoint",
            "adverse event",
            "safety",
            "efficacy",
            "biostatistics",
            "regulatory",
            "gcp",
            "ich",
            "ind",
            "nda",
            "cro",
            "monitoring",
            "informed consent",
            "patient recruitment",
            "statistical analysis",
            "pharmacokinetics",
        ]
        await asyncio.sleep(1)
        combined_text = message.lower()
        clinical_matches = sum(1 for keyword in clinical_keywords if keyword in combined_text)

        result = "CLINICAL_TRIAL_RELATED" if clinical_matches > 0 else "NON_CLINICAL_RELATED"
        print(f"Classification result: {result} (found {clinical_matches} matches)")
        return result


def create_crew():
    # Initialize LLM with error handling
    try:
        ollama_llm = LLM(model="ollama/phi4-mini:latest", base_url="http://localhost:11434", api_version="ollama")
        print("✓ LLM initialized successfully")
    except Exception as e:
        print(f"✗ Failed to initialize LLM: {e}")
        raise

    # Define agent with simplified configuration
    topic_router = Agent(
        role="Topic Classification Specialist",
        goal="Classify messages as clinical trial related or not",
        backstory="You are an expert at identifying clinical trial related content.",
        verbose=True,
        memory=False,
        allow_delegation=False,
        llm=ollama_llm,
        max_iter=2,  # Reduced iterations
    )

    # Define task with clearer description
    classification_task = Task(
        description="""
        You are tasked with classifying messages related to clinical trials.
        You will receive a message and must determine if it is related to clinical trials or not.
        classify this message: {message}
        """,
        expected_output="Classification result: either CLINICAL_TRIAL_RELATED or NON_CLINICAL_TRIAL_RELATED",
        agent=topic_router,
    )

    # Create crew
    crew = Crew(
        agents=[topic_router],
        tasks=[classification_task],
        process=Process.sequential,
        verbose=True,
        memory=False,  # Disable memory to avoid potential issues
    )

    return crew


def main():

    try:
        crew = create_crew()
        print("✓ Crew created successfully")

        print("\nExecuting classification task...")
        response = crew.kickoff(
            inputs={
                "message": "The Future of AI in Healthcare",
            }
        )
        print(f"\nFinal Response: {response}")

    except Exception as e:
        print(f"✗ Error during execution: {e}")
        import traceback

        traceback.print_exc()


if __name__ == "__main__":
    main()
