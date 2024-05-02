import { createSlice } from "@reduxjs/toolkit";

export const headerSlice = createSlice({
  name: "header",
  initialState: {
    selected: "", // Selected menu
  },
  reducers: {
    setMenu: (state, action) => {
      state.selected = action.payload;
    },
  },
});

export const { setMenu } = headerSlice.actions;
export default headerSlice.reducer;