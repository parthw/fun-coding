const data = [
  { rack_type: "A" },
  { rack_type: "B" },
  { rack_type: "A" },
  { rack_type: "B" },
];

const svg = document.getElementById("mySvg");
const domContainer = document.getElementById("dom-container");

data.forEach((item, index) => {
  const rect = document.createElementNS("http://www.w3.org/2000/svg", "rect");
  rect.setAttribute("x", 20 + index * 100);
  rect.setAttribute("y", 50);
  rect.setAttribute("width", 80);
  rect.setAttribute("height", 100);
  rect.setAttribute("class", `${item.rack_type}-svg`);
  rect.setAttribute("data-rack-type", item.rack_type);
  svg.appendChild(rect);
});

const colors = ["red", "green", "blue", "purple"];
const assignedColors = {};
data.forEach((item, index) => {
  if (!assignedColors[item.rack_type]) {
    assignedColors[item.rack_type] = colors[index % colors.length];

    const style = document.createElement("style");
    style.textContent = `
    .${item.rack_type}-svg {
      fill: ${assignedColors[item.rack_type]}; 
      stroke: black;
      stroke-width: 2px;
    }
  `;
    document.head.appendChild(style);
  }
});

const distinctRackTypes = [...new Set(data.map((item) => item.rack_type))];
const buttonContainer = document.getElementById("index-buttons");
distinctRackTypes.forEach((rackType) => {
  const button = document.createElement("button");
  button.textContent = `Show ${rackType}`;
  button.classList.add("rack-button");
  button.setAttribute("data-rack", rackType);

  button.addEventListener("click", (e) => {
    // Highlight the clicked button's rack type and grey out others
    document.querySelectorAll(".rack-button").forEach((btn) => {
      if (btn === e.target) {
        btn.classList.add("highlight");
      } else {
        btn.classList.remove("highlight");
      }
    });

    document.querySelectorAll("rect").forEach((rect) => {
      if (rect.getAttribute("data-rack-type") === rackType) {
        rect.style.fill = assignedColors[rackType]; // Show original color
      } else {
        rect.style.fill = "grey"; // Grey out other colors
      }
    });
  });

  buttonContainer.appendChild(button);
});

// Reset all colors when clicking outside the button container
document.addEventListener("click", (e) => {
  if (!buttonContainer.contains(e.target)) {
    document.querySelectorAll("rect").forEach((rect) => {
      const rackType = rect.getAttribute("data-rack-type");
      rect.style.fill = assignedColors[rackType]; // Reset to original color
    });

    document.querySelectorAll(".rack-button").forEach((btn) => {
      btn.classList.remove("highlight");
    });
  }
});
