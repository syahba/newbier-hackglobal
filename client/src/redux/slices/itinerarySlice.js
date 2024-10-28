import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const itinerarySlice = createSlice({
  name: 'chat',
  initialState: {
    itineraries: [],
    itineraryMessage: ''
  },
  reducers: {
    setItinerary(state, action) {
      state.chats = action.payload;
    },
    setItineraryMessage(state, action) {
      state.chatMessage = action.payload
    }
  }
});

export const { setItinerary, setItineraryMessage } = itinerarySlice.actions;

export const getItineraries = () => {
  return async (dispatch) => {
    const { data } = await axios({
      method: 'post',
      url: `http://localhost:8000/api/itinerary`,
      data: payload
    });

    dispatch(setItinerary(data))
  };
};

export const sendChat = (payload) => {
  return async (dispatch) => {
    const { message } = await axios({
      method: 'post',
      url: `http://localhost:8000/api/chat`,
      data: payload
    });

    dispatch(setChatMessage(message))
  };
};

export default itinerarySlice.reducer;