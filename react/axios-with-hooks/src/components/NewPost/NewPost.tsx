import React from "react";
import styles from "./NewPost.module.css";

type NewPostPropsType = {};

export default function NewPost(props: NewPostPropsType) {
  return (
    <div className={styles.NewPost}>
      <h1>Add a Post</h1>
      <label>Title</label>
      <input type="text" />
      <label>Content</label>
      <textarea />
      <label>Author</label>
      <select>
        <option value="Gin">Gin</option>
        <option value="Shin">Shin</option>
        <option value="Kagura">Kagura</option>
      </select>
      <button>Add Post</button>
    </div>
  );
}
