import React from "react";

export const appToPersonContext = React.createContext({
  deleteHandler: (index: number) => {},
  changeTextHandler: (
    index: number,
    event: React.ChangeEvent<HTMLInputElement>
  ) => {},
});
