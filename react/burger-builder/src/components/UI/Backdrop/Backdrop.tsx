import React from "react";
import styles from "./Backdrop.module.css";

export default function Backdrop(props: {
  show: boolean;
  clicked: React.MouseEventHandler;
}) {
  return props.show ? (
    <div className={styles.Backdrop} onClick={props.clicked}></div>
  ) : null;
}
