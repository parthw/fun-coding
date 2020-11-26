import React, { useState, useEffect } from "react";
import AppClasses from "./App.module.css";
import Persons from "../Components/Persons/Persons";
import Counter from "../Components/Counter/Counter";
import { appToPersonContext } from "../Contexts/appToPerson";

function App() {
  useEffect(() => {
    console.log("App useEffect");
  });
  let [buttonClassState, setButtonClassState] = useState({
    class: [AppClasses.App, AppClasses.green],
    toogleStatus: 0,
  });
  const toogleButtonHandler = () => {
    if (!buttonClassState.toogleStatus) {
      setButtonClassState({
        class: [AppClasses.App, AppClasses.red],
        toogleStatus: 1,
      });
    } else {
      setButtonClassState({
        class: [AppClasses.App, AppClasses.green],
        toogleStatus: 0,
      });
    }
  };

  let [namesState, setNamesState] = useState(["Noodles", "Pizza", "Pasta"]);
  const deleteParaHandler = (index: number) => {
    const existingState = [...namesState];
    existingState.splice(index, 1);
    setNamesState(existingState);
  };

  let personsComponents: JSX.Element | string;
  if (!buttonClassState.toogleStatus) {
    personsComponents = <Persons names={namesState} />;
  } else {
    personsComponents = "";
  }

  const changeTextHandler = (
    index: number,
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const newNameState = [...namesState];
    newNameState[index] = event.target.value;
    setNamesState(newNameState);
  };

  return (
    <div className={AppClasses.App}>
      <div>
        <button
          className={buttonClassState.class.join(" ")}
          onClick={toogleButtonHandler}
        >
          Toogle Button
        </button>
      </div>
      <div>
        <Counter names={namesState} />
        <appToPersonContext.Provider
          value={{
            deleteHandler: deleteParaHandler,
            changeTextHandler: changeTextHandler,
          }}
        >
          {personsComponents}
        </appToPersonContext.Provider>
      </div>
    </div>
  );
}

export default App;
