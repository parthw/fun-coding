import React from "react";
import styles from "./Posts.module.css";
import { Link, RouteComponentProps } from "react-router-dom";

export type postType = {
  title: string;
  author: string;
  id: string;
};

type PostsPropsType = {
  postsData: postType[];
  postClick: Function;
} & RouteComponentProps;

export default function Posts(props: PostsPropsType) {
  const posts = props.postsData.map((post: postType, index: number) => {
    return (
      <Link to={`${props.match.url}/${post.id}`} key={post.id}>
        <div
          className={styles.Post}
          key={index}
          onClick={() => props.postClick(post.id)}
        >
          <h3>
            <strong>{post.title}</strong>
          </h3>
          <div className={styles.Author}>{post.author}</div>
        </div>
      </Link>
    );
  });
  return <div className={styles.Posts}>{posts}</div>;
}
