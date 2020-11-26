import React from "react";
import styles from "./NavigationItems.module.css";

import NavigationItem from "./NavigationItem/NavigationItem";

export default function NavigationItems(props: {}) {
  const active = true;
  return (
    <ul className={styles.NavigationItems}>
      <NavigationItem link="/" active={active}>
        Burger Builder
      </NavigationItem>
      <NavigationItem link="/" active={!active}>
        Checkout
      </NavigationItem>
    </ul>
  );
}
