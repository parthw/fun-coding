import React from "react";
import styles from "./Button.module.css";

type ButtonPropsType = {
  btnType: "Success" | "Danger";
  click: React.MouseEventHandler;
};

export default function Button(
  props: React.PropsWithChildren<ButtonPropsType>
) {
  return (
    <button
      className={[styles.Button, styles[props.btnType]].join(" ")}
      onClick={props.click}
    >
      {props.children}
    </button>
  );
}
