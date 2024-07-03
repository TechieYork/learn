import { createSlice } from "@reduxjs/toolkit";

export const courseLearningSlice = createSlice({
  name: "courseLearning",
  initialState: {
    itemKey: "0-0-1",
    section: "",
    topic: "",
  },
  reducers: {
    setItemKey: (state, action) => {
      state.itemKey = action.payload;
    },
    setSection: (state, action) => {
      state.section = action.payload;
    },
    setTopic: (state, action) => {
      state.topic = action.payload;
    },
  },
});

export const {
  setItemKey,
  setSection,
  setTopic,
} = courseLearningSlice.actions;
export default courseLearningSlice.reducer;
