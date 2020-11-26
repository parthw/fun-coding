import React from "react";
import styles from "./Spinner.module.css";

type SpinnerPropsType = {};

export default function Spinner(props: SpinnerPropsType) {
  return <div className={styles.Loader}>Loading...</div>;
}
