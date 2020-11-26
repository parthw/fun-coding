import React from "react";
import styles from "./Toolbar.module.css";

import Logo from "../../Logo/Logo";
import NavigationItems from "../NavigationItems/NavigationItems";

export default function Toolbar(props: {
  toogleSideDrawer: React.MouseEventHandler;
}) {
  return (
    <header className={styles.Toolbar}>
      <div onClick={props.toogleSideDrawer}>MENU</div>
      <div className={styles.ResponsiveLogo}>
        <Logo />
      </div>
      <nav className={styles.DesktopOnly}>
        <NavigationItems />
      </nav>
    </header>
  );
}
