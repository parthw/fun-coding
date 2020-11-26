import React, { useState } from "react";
import styles from "./Layout.module.css";

import Toolbar from "../Navigation/Toolbar/Toolbar";
import SideDrawer from "../Navigation/SideDrawer/SideDrawer";

type LayoutProps = {};

export default function Layout(props: React.PropsWithChildren<LayoutProps>) {
  let [showSideDrawerState, setShowSideDrawerState] = useState<boolean>(true);
  const toogleSideDrawerHandler = () => {
    setShowSideDrawerState(!showSideDrawerState);
  };
  return (
    <React.Fragment>
      <Toolbar toogleSideDrawer={toogleSideDrawerHandler} />
      <SideDrawer
        show={showSideDrawerState}
        backdropClicked={toogleSideDrawerHandler}
      />
      <main className={styles.Content}>{props.children}</main>
    </React.Fragment>
  );
}
