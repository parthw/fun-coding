import React, { useState, useEffect, useRef } from "react";
import CheckoutSummary from "../../components/Order/CheckoutSummary/CheckoutSummary";
import { RouteComponentProps, Route } from "react-router-dom";
import ContactData from "./ContactData/ContactData";

export default function Checkout(props: RouteComponentProps) {
  let [ingredientsState, setIngredientsState] = useState<{
    [key: string]: number;
  }>({});

  let [totalPriceState, setTotalPriceState] = useState(0);

  let urlSearchRef = useRef<string>(props.location.search);
  useEffect(() => {
    const query = new URLSearchParams(urlSearchRef.current);
    const ingredients: { [key: string]: number } = {};
    for (let [key, value] of query.entries()) {
      if (key === "price") {
        totalPriceState = +value;
      } else {
        ingredients[key] = +value;
      }
    }
    setIngredientsState(ingredients);
  }, []);

  const checkoutCancelledHandler = () => {
    props.history.goBack();
  };
  const checkoutContinuedHandler = () => {
    props.history.replace(props.match.url + "/contact-data");
  };

  return (
    <div>
      <CheckoutSummary
        ingredients={ingredientsState}
        checkoutCancelled={checkoutCancelledHandler}
        checkoutContinued={checkoutContinuedHandler}
      />

      <Route
        path={props.match.path + "/contact-data"}
        render={(routerProps) => (
          <ContactData
            {...routerProps}
            ingredients={ingredientsState}
            totalPrice={totalPriceState}
          />
        )}
      />
    </div>
  );
}
