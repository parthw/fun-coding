import React, { useState, useEffect } from "react";
import Modal from "../../components/UI/Modal/Modal";
import { AxiosInstance } from "../../api/axios";
import { RouteComponentProps } from "react-router-dom";

const withErrorHandler = (
  WrappedComponent: React.ComponentType<RouteComponentProps>,
  axios: AxiosInstance
) => {
  return (props: RouteComponentProps) => {
    let [errorState, setErrorState] = useState<Error | null>(null);
    let [errorStatusState, setErrorStatusState] = useState<boolean>(false);

    useEffect(() => {
      let reqInterceptor = axios.interceptors.request.use((req) => {
        setErrorState(null);
        setErrorStatusState(false);
        return req;
      });
      let respInterceptor = axios.interceptors.response.use(
        (resp) => resp,
        (err) => {
          setErrorState(err);
          setErrorStatusState(true);
        }
      );

      return () => {
        axios.interceptors.request.eject(reqInterceptor);
        axios.interceptors.response.eject(respInterceptor);
      };
    }, []);

    const errorConfirmedHandler = () => {
      setErrorState(null);
      setErrorStatusState(false);
    };

    return (
      <React.Fragment>
        <Modal show={errorStatusState} modalClicked={errorConfirmedHandler}>
          Something didn't work!
          {errorState?.message}
        </Modal>
        <WrappedComponent {...props} />
      </React.Fragment>
    );
  };
};

export default withErrorHandler;
