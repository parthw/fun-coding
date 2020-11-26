import React from "react";
import styles from "./FullPost.module.css";

export type fullPostType = {
  id: string;
  body: string;
  title: string;
};

type FullPostPropsType = {
  fullPost: fullPostType;
  deleteHandler: Function;
};

export default function FullPost(props: FullPostPropsType) {
  return (
    <div className={styles.FullPost}>
      <h2>{props.fullPost.title}</h2>
      <div>{props.fullPost.body}</div>
      <button onClick={() => props.deleteHandler(props.fullPost.id)}>
        Delete
      </button>
    </div>
  );
}
