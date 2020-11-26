import React, { useEffect } from "react";
import Button from "../../UI/Button/Button";

type OrderSummaryPropsType = {
  ingredientsWithCount: { [key: string]: number };
  purchaseCanceled: React.MouseEventHandler;
  purchaseContinued: React.MouseEventHandler;
  totalPrice: number;
};

export default function OrderSummary(props: OrderSummaryPropsType) {
  useEffect(() => {
    console.log("Order Summary");
  });
  const orderSummary = Object.keys(props.ingredientsWithCount).map((ingre) => (
    <li key={ingre}>
      {ingre}: {props.ingredientsWithCount[ingre]}
    </li>
  ));
  return (
    <div>
      <h3>Order Summary</h3>
      <ul>{orderSummary}</ul>
      <p>Continue to Checkout ?</p>
      <p>
        <strong> Total price: {props.totalPrice.toFixed(2)} </strong>
      </p>
      <Button btnType="Danger" click={props.purchaseCanceled}>
        Cancel
      </Button>
      <Button btnType="Success" click={props.purchaseContinued}>
        Continue
      </Button>
    </div>
  );
}
