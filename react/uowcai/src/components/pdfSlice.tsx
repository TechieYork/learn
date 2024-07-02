import { createSlice } from "@reduxjs/toolkit";

export const pdfSlice = createSlice({
  name: "pdf",
  initialState: {
    loading: true,
  },
  reducers: {
    setPdfLoading: (state, action) => {
      console.log("setPdfLoading", action.payload);
      state.loading = action.payload;
    }
  },
});

export const { setPdfLoading } = pdfSlice.actions;
export default pdfSlice.reducer;