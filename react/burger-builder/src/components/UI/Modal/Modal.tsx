import React, { useEffect } from "react";
import styles from "./Modal.module.css";

import Backdrop from "../Backdrop/Backdrop";

type ModalProps = {
  show: boolean;
  modalClicked: React.MouseEventHandler;
};

function Modal(props: React.PropsWithChildren<ModalProps>) {
  useEffect(() => {
    console.log("Modal");
  });
  return (
    <React.Fragment>
      <Backdrop show={props.show} clicked={props.modalClicked} />
      <div
        className={styles.Modal}
        style={{
          transform: props.show ? "translateY(0)" : "translateY(-100vh)",
          opacity: props.show ? "1" : "0",
        }}
      >
        {props.children}
      </div>
    </React.Fragment>
  );
}

export default React.memo(
  Modal,
  (prevProps, nextProps) =>
    prevProps.show === nextProps.show &&
    prevProps.children === nextProps.children
);
