import { createSlice } from "@reduxjs/toolkit";

export const aboutSlice = createSlice({
  name: "about",
  initialState: {
    about: "About",
  },
  reducers: {},
});

export default aboutSlice.reducer;