import React from "react";
import BurgerIngredient from "./BurgerIngredients/BurgerIngredients";
import styles from "./Burger.module.css";

type ingredientsType = {
  [ingredient: string]: number;
};

export default function Burger(props: { ingredients: ingredientsType }) {
  let transformedIngredients = Object.keys(props.ingredients)
    .map((igKey) =>
      [...Array(props.ingredients[igKey])].map((_, i) => (
        <BurgerIngredient type={igKey} key={igKey + i} />
      ))
    )
    .reduce((pv, curr) => pv.concat(curr), []);

  if (transformedIngredients.length === 0) {
    transformedIngredients = [
      <div key="00">Please start adding some items!</div>,
    ];
  }
  return (
    <div className={styles.Burger}>
      <BurgerIngredient type="bread-top" />
      {transformedIngredients}
      <BurgerIngredient type="bread-bottom" />
    </div>
  );
}
