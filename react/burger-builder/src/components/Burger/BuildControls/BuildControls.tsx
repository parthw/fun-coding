import React from "react";
import styles from "./BuildControls.module.css";

import BuildControl from "./BuildControl/BuildControl";

type BuildControlsProps = {
  addIngredient: Function;
  removeIngredient: Function;
  disableRemoveButton: { [key: string]: boolean };
  price: number;
  purchasable: boolean;
  ordered: React.MouseEventHandler;
};

const controls = [
  { label: "Salad", type: "salad" },
  { label: "Cheese", type: "cheese" },
  { label: "Bacon", type: "bacon" },
  { label: "Meat", type: "meat" },
];
export default function BuildControls(props: BuildControlsProps) {
  return (
    <div className={styles.BuildControls}>
      <p>
        Current Prince: <strong>{props.price.toFixed(2)}</strong>
      </p>

      {controls.map((ctrl) => (
        <BuildControl
          key={ctrl.label}
          label={ctrl.label}
          addIngredient={() => props.addIngredient(ctrl.type)}
          removeIngredient={() => props.removeIngredient(ctrl.type)}
          disableButton={props.disableRemoveButton[ctrl.type]}
        />
      ))}

      <button
        className={styles.OrderButton}
        disabled={props.purchasable}
        onClick={props.ordered}
      >
        ORDER NOW
      </button>
    </div>
  );
}
