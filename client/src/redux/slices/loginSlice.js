import { createSlice } from "@reduxjs/toolkit";

const loginSlice = createSlice({
  name: 'login',
  initialState: {
    name: ''
  },
  reducers: {
    setUser(state, action) {
      state.name = action.payload;
    }
  }
});

const { setUser } = loginSlice.actions;

export const login = (payload) => {
  return setUser(payload);
}

export default loginSlice.reducer;