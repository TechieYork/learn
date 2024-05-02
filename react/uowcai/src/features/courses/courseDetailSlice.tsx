import { createSlice } from "@reduxjs/toolkit";

export const courseDetailSlice = createSlice({
  name: "courses",
  initialState: {
    courseDetailImageLoaded: false,
  },
  reducers: {
    setCourseDetailImageLoaded: (state, action) => {
      state.courseDetailImageLoaded = action.payload;
    },
  },
});

export const { setCourseDetailImageLoaded } = courseDetailSlice.actions;
export default courseDetailSlice.reducer;
