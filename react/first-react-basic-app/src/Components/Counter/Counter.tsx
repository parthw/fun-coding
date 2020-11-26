import React, { useEffect } from "react";

export default function Counter(props: { names: string[] }) {
  useEffect(() => {
    console.log("Counter useEffect");
  });
  return <div>Total Count is {props.names.length}</div>;
}
