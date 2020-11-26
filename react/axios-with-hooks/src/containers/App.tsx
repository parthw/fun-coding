import React, { useState, useEffect, Suspense } from "react";
import styles from "./App.module.css";

import Posts from "../components/Posts/Posts";
import { postType } from "../components/Posts/Posts";

import { fullPostType } from "../components/FullPost/FullPost";
import NewPost from "../components/NewPost/NewPost";
import { BrowserRouter, Link, Route, Switch } from "react-router-dom";

import jpaxios from "../apis/jsonplaceholder.axios";

const FullPost = React.lazy(() => import("../components/FullPost/FullPost"));

export default function App(props: {}) {
  let [postsState, setPostsState] = useState<postType[]>();
  let [fullPostState, setFullPostState] = useState<fullPostType>({
    id: "0",
    body: "body",
    title: "Title",
  });

  useEffect(() => {
    jpaxios.get<postType[]>("/posts").then((response) => {
      const postsData = response.data.slice(0, 3);
      const posts = postsData.map((post: postType) => {
        return { title: post.title, author: "katsura", id: post.id };
      });
      setPostsState(posts);
    });
  }, []);

  const fullPostButtonHandler = (id: string) => {
    fullPostState?.id !== id &&
      jpaxios.get<fullPostType>(`/posts/${id}`).then((response) => {
        setFullPostState({
          id: id,
          body: response.data.body,
          title: response.data.title,
        });
      });
  };

  const fullPostDeleteHandler = (id: string) => {
    id !== "0" &&
      setFullPostState({
        id: "0",
        body: "body",
        title: "title",
      });
  };

  return (
    <BrowserRouter>
      <div className={styles.APP}>
        <ul>
          <li>
            <Link to="/new-post">New Post</Link>
          </li>
          <li>
            <Link to="/posts">Posts</Link>
          </li>
          <li>
            <Link to="/full-post">Full Post</Link>
          </li>
        </ul>

        <Switch>
          <Route
            path="/posts"
            exact
            render={(routeProps) =>
              postsState ? (
                <Posts
                  {...routeProps}
                  postsData={postsState}
                  postClick={fullPostButtonHandler}
                />
              ) : (
                <div>Loading...</div>
              )
            }
          />
          <Route path="/new-post" exact component={NewPost} />
        </Switch>
        <Route
          path={["/full-post", "/posts/:id"]}
          exact
          render={() => (
            <Suspense fallback={<h2>Loading...</h2>}>
              <FullPost
                fullPost={fullPostState}
                deleteHandler={fullPostDeleteHandler}
              />
            </Suspense>
          )}
        />
      </div>
    </BrowserRouter>
  );
}
