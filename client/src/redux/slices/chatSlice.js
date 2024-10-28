import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const chatSlice = createSlice({
  name: 'chat',
  initialState: {
    chats: [],
    chatMessage: ''
  },
  reducers: {
    setChat(state, action) {
      state.chats = action.payload;
    },
    setChatMessage(state, action) {
      state.chatMessage = action.payload
    }
  }
});

export const { setChat, setChatMessage } = chatSlice.actions;

export const getChats = () => {
  return async (dispatch) => {
    const { data } = await axios({
      method: 'get',
      url: `${process.env.HOST}/api/chat`,
      data: payload
    });

    dispatch(setChat(data))
  };
};

export const sendChat = (payload) => {
  return async (dispatch) => {
    const { message } = await axios({
      method: 'post',
      url: `${process.env.HOST}/api/chat`,
      data: payload
    });

    dispatch(setChatMessage(message))
  };
};

export default chatSlice.reducer;