import { createSlice } from "@reduxjs/toolkit";

export const homeSlice = createSlice({
  name: "home",
  initialState: {
    home: "home",
  },
  reducers: {},
});

export default homeSlice.reducer;