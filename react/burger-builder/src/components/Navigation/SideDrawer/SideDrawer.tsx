import React from "react";
import styles from "./SideDrawer.module.css";

import Logo from "../../Logo/Logo";
import NavigationItems from "../NavigationItems/NavigationItems";
import Backdrop from "../../UI/Backdrop/Backdrop";

export default function SideDrawer(props: {
  show: boolean;
  backdropClicked: React.MouseEventHandler;
}) {
  let sideDrawerTooglingStyle = [];
  if (!props.show) {
    sideDrawerTooglingStyle = [styles.SideDrawer, styles.Close];
  } else {
    sideDrawerTooglingStyle = [styles.SideDrawer, styles.Open];
  }
  return (
    <React.Fragment>
      <Backdrop show={props.show} clicked={props.backdropClicked} />
      <div className={sideDrawerTooglingStyle.join(" ")}>
        <div className={styles.ResponsiveLogo}>
          <Logo />
        </div>
        <nav>
          <NavigationItems />
        </nav>
      </div>
    </React.Fragment>
  );
}
