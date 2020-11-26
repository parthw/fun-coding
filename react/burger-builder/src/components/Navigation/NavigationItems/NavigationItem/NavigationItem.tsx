import React from "react";
import styles from "./NavigationItem.module.css";

type NavigationItemPropsType = { link: string; active: boolean };

export default function NavigationItem(
  props: React.PropsWithChildren<NavigationItemPropsType>
) {
  return (
    <li className={styles.NavigationItem}>
      <a href={props.link} className={props.active ? styles.active : undefined}>
        {props.children}
      </a>
    </li>
  );
}
