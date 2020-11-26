import React from "react";
import styles from "./BuildControl.module.css";

type BuildControlType = {
  label: string;
  addIngredient: React.MouseEventHandler;
  removeIngredient: React.MouseEventHandler;
  disableButton: boolean;
};

export default function BuildControl(props: BuildControlType) {
  return (
    <div className={styles.BuildControl}>
      <div className={styles.Label}>{props.label}</div>
      <button
        className={styles.Less}
        onClick={props.removeIngredient}
        disabled={props.disableButton}
      >
        Less
      </button>
      <button className={styles.More} onClick={props.addIngredient}>
        More
      </button>
    </div>
  );
}
