import React, { useEffect } from "react";
import Person from "./Person/Person";

export default function Persons(props: { names: string[] }) {
  useEffect(() => {
    console.log("Persons useEffect");
  });
  return (
    <div>
      {props.names.map((name, index) => {
        return <Person name={name} key={index} index={index} />;
      })}
    </div>
  );
}
