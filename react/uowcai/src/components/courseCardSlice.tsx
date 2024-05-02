import { createSlice } from "@reduxjs/toolkit";

export const courseCardSlice = createSlice({
  name: "courseCard",
  initialState: {
    imageLoaded: false,
  },
  reducers: {
    setImageLoaded: (state, action) => {
      state.imageLoaded = action.payload;
    }
  },
});

export const { setImageLoaded } = courseCardSlice.actions;
export default courseCardSlice.reducer;