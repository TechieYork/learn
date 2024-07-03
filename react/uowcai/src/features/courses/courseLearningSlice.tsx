import { createSlice } from "@reduxjs/toolkit";

export const courseLearningSlice = createSlice({
  name: "courseLearning",
  initialState: {
    topic: "0-0-1",
  },
  reducers: {
    setTopic: (state, action) => {
      state.topic = action.payload;
    },
  },
});

export const { setTopic } = courseLearningSlice.actions;
export default courseLearningSlice.reducer;
