import React, { MouseEventHandler } from "react";
import styles from "./CheckoutSummary.module.css";
import Button from "../../UI/Button/Button";
import Burger from "../../Burger/Burger";

type CheckoutSummaryPropsType = {
  ingredients: { [key: string]: number };
  checkoutContinued: MouseEventHandler;
  checkoutCancelled: MouseEventHandler;
};

export default function CheckoutSummary(props: CheckoutSummaryPropsType) {
  return (
    <div className={styles.CheckoutSummary}>
      <h1>We hope it tastes well!</h1>
      <div style={{ width: "100%", margin: "auto" }}>
        <Burger ingredients={props.ingredients} />
      </div>
      <Button btnType="Danger" click={props.checkoutCancelled}>
        Cancel
      </Button>
      <Button btnType="Success" click={props.checkoutContinued}>
        Continue
      </Button>
    </div>
  );
}
