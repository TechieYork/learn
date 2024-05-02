import { createSlice } from "@reduxjs/toolkit";

export const coursesSlice = createSlice({
  name: "courses",
  initialState: {
    coursesImageLoaded: false,
  },
  reducers: {
    setCoursesImageLoaded: (state, action) => {
      state.coursesImageLoaded = action.payload;
    },
  },
});

export const { setCoursesImageLoaded } = coursesSlice.actions;
export default coursesSlice.reducer;
