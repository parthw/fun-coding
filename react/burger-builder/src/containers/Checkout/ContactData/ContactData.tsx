import React, { useState } from "react";
import styles from "./ContactData.module.css";
import Button from "../../../components/UI/Button/Button";
import { RouteComponentProps } from "react-router-dom";
import firebaseAxios from "../../../api/firebase.axios";

type ContactDataType = {
  ingredients: { [key: string]: number };
  totalPrice: number;
} & RouteComponentProps;

export default function ContactData(props: ContactDataType) {
  let [userDetails, setUserDetailsState] = useState({
    name: "Name",
    email: "Email@email.com",
    address: "address",
  });

  let [loadingState, setLoadingState] = useState<boolean>(false);

  const orderHandler = (event: React.MouseEvent) => {
    event.preventDefault();
    setLoadingState(true);
    const order = {
      totalPrice: props.totalPrice,
      ingredients: props.ingredients,
      customer: {
        name: "Parth Wadhwa",
        address: "Address",
        email: "email@email.com",
      },
    };
    firebaseAxios
      .post("/orders.json", order)
      .then((response) => {
        setLoadingState(false);
      })
      .catch((err) => {
        console.log(err);
        setLoadingState(false);
      });
  };

  return (
    <div className={styles.ContactData}>
      <h4>Enter your contact data!</h4>
      <form>
        <input type="text" name="name" placeholder="Your Name" />
        <input type="email" name="email" placeholder="Your Email" />
        <input type="text" name="address" placeholder="Your Address" />
        <Button btnType="Success" click={orderHandler}>
          done
        </Button>
      </form>
    </div>
  );
}
