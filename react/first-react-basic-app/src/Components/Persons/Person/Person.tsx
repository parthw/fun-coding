import React, { useContext, useEffect, useRef } from "react";
import PersonClasses from "./Person.module.css";
import { appToPersonContext } from "../../../Contexts/appToPerson";

export default function Person(props: { name: string; index: number }) {
  const { deleteHandler, changeTextHandler } = useContext(appToPersonContext);
  const inputRef = useRef<HTMLInputElement>(null);
  useEffect(() => {
    console.log("Person useEffect");
    if (inputRef && inputRef.current && props.index === 1) {
      inputRef.current.focus();
    }
    return () => {
      console.log("Person useEffect Return");
    };
  });
  return (
    <div>
      <p
        className={PersonClasses.Person}
        onClick={() => deleteHandler(props.index)}
      >
        Hi! My name is {props.name}
        <br />
      </p>
      <input
        type="text"
        ref={inputRef}
        onChange={(event) => changeTextHandler(props.index, event)}
      />
    </div>
  );
}
