import React from "react";
import PropTypes from "prop-types";
import styles from "./BurgerIngredients.module.css";

export default function BurgerIndegridient(props: { type: string }) {
  let indegridient: JSX.Element;
  switch (props.type) {
    case "bread-bottom":
      indegridient = <div className={styles.BreadBottom}></div>;
      break;
    case "bread-top":
      indegridient = (
        <div className={styles.BreadTop}>
          <div className={styles.Seeds1}></div>
          <div className={styles.Seeds2}></div>
        </div>
      );
      break;
    case "meat":
      indegridient = <div className={styles.Meat}></div>;
      break;
    case "cheese":
      indegridient = <div className={styles.Cheese}></div>;
      break;
    case "bacon":
      indegridient = <div className={styles.Bacon}></div>;
      break;
    case "salad":
      indegridient = <div className={styles.Salad}></div>;
      break;
    default:
      indegridient = <div></div>;
  }

  return indegridient;
}

BurgerIndegridient.propTypes = {
  type: PropTypes.string.isRequired,
};
