import { configureStore } from "@reduxjs/toolkit";
import login from "./slices/loginSlice";
import chat from "./slices/chatSlice";

const store = configureStore({
  reducer: {
    login,
    chat
  }
});

export default store;