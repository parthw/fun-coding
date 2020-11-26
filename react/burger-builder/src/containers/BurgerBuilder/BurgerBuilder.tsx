import React, { useState, useEffect } from "react";
import Burger from "../../components/Burger/Burger";
import BuildControls from "../../components/Burger/BuildControls/BuildControls";
import Modal from "../../components/UI/Modal/Modal";
import OrderSummary from "../../components/Burger/OrderSummary/OrderSummary";
import Spinner from "../../components/UI/Spinner/Spinner";

import withErrorHandler from "../withErrorHandler/withErrorHandler";
import faxios from "../../api/firebase.axios";
import { RouteComponentProps } from "react-router-dom";

type ingredientType = {
  [key: string]: number;
};

const INGREDIENTS_PRICES: ingredientType = {
  salad: 0.5,
  cheese: 0.2,
  bacon: 1.5,
  meat: 0.9,
};

function BurgerBuilder(props: RouteComponentProps) {
  let [ingredientsState, setIngredientsState] = useState<ingredientType | null>(
    null
  );
  let [ingredientsTotalPrice, setIngredientsTotalPrice] = useState<number>(4);
  let [purchasingState, setPurchasingState] = useState<boolean>(false);
  let [errorStatusState, setErrorStatusState] = useState<boolean>(false);

  useEffect(() => {
    faxios
      .get("/ingredients.json")
      .then((resp) => {
        setIngredientsState(resp.data);
        setErrorStatusState(false);
      })
      .catch((err) => {
        setErrorStatusState(true);
      });
  }, []);

  const purchasable = () => {
    if (!ingredientsState) {
      return false;
    }
    const totalIngredientsPriceSum = Object.values(ingredientsState).reduce(
      (sum, el) => sum + el,
      0
    );
    if (totalIngredientsPriceSum === 0) return false;
    return true;
  };

  const addIngredientHandler = (type: string) => {
    if (!ingredientsState) {
      return;
    }
    let updatedIngredient = {
      ...ingredientsState,
    };
    updatedIngredient[type] += 1;

    setIngredientsState(updatedIngredient);
    let updatedIngredientTotalPrice =
      ingredientsTotalPrice + INGREDIENTS_PRICES[type];
    setIngredientsTotalPrice(updatedIngredientTotalPrice);
  };

  const removeIngredientHandler = (type: string) => {
    if (!ingredientsState) {
      return;
    }
    let updatedIngredient = {
      ...ingredientsState,
    };
    if (updatedIngredient[type] === 0) {
      return;
    }
    updatedIngredient[type] -= 1;

    setIngredientsState(updatedIngredient);
    let updatedIngredientTotalPrice =
      ingredientsTotalPrice - INGREDIENTS_PRICES[type];
    setIngredientsTotalPrice(updatedIngredientTotalPrice);
  };

  const purchaseHandler = () => {
    setPurchasingState(true);
  };

  const purchaseCancelHandler = () => {
    setPurchasingState(false);
  };

  const purchaseContinueHandler = () => {
    const queryParams = [];
    if (!ingredientsState) {
      return;
    }
    for (let i of Object.keys(ingredientsState)) {
      queryParams.push(
        encodeURIComponent(i) + "=" + encodeURIComponent(ingredientsState[i])
      );
    }
    queryParams.push("price=" + ingredientsTotalPrice);
    const queryString = queryParams.join("&");
    console.log(`Printing Query String ${queryString}`);
    props.history.push({
      pathname: "/checkout",
      search: "?" + queryString,
    });
  };

  const disableRemoveButton: { [key: string]: boolean } = {};
  let orderSummaryElement = <Spinner />;
  let burgerElement = errorStatusState ? (
    <p>Ingredients can't be loaded!</p>
  ) : (
    <Spinner />
  );

  if (ingredientsState) {
    for (let ingre of Object.keys(ingredientsState)) {
      if (ingredientsState[ingre] === 0) {
        disableRemoveButton[ingre] = true;
      } else {
        disableRemoveButton[ingre] = false;
      }
    }

    burgerElement = <Burger ingredients={ingredientsState} />;
    orderSummaryElement = (
      <OrderSummary
        ingredientsWithCount={ingredientsState}
        purchaseCanceled={purchaseCancelHandler}
        purchaseContinued={purchaseContinueHandler}
        totalPrice={ingredientsTotalPrice}
      ></OrderSummary>
    );
  }

  return (
    <React.Fragment>
      {burgerElement}
      <Modal show={purchasingState} modalClicked={purchaseCancelHandler}>
        {orderSummaryElement}
      </Modal>
      <BuildControls
        disableRemoveButton={disableRemoveButton}
        addIngredient={addIngredientHandler}
        removeIngredient={removeIngredientHandler}
        price={ingredientsTotalPrice}
        purchasable={!purchasable()}
        ordered={purchaseHandler}
      />
    </React.Fragment>
  );
}

export default withErrorHandler(BurgerBuilder, faxios);
