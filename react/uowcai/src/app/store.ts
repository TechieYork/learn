import { configureStore } from '@reduxjs/toolkit'
import aboutReducer from '../features/about/aboutSlice'
import homeReducer from '../features/home/homeSlice'
import coursesReducer from '../features/courses/coursesSlice'
import courseDetailReducer from '../features/courses/courseDetailSlice'
import certificationReducer from '../features/certification/certificationSlice'
import notFoundReducer from '../features/not-found/notFoundSlice'
import courseCardReducer from '../components/courseCardSlice'
import headerReducer from '../components/headerSlice'
import footerReducer from '../components/footerSlice'
import pdfReducer from '../components/pdfSlice'

export default configureStore({
    reducer: {
        home: homeReducer,
        courses: coursesReducer,
        certification: certificationReducer,
        about: aboutReducer,
        notFound: notFoundReducer,
        courseCard: courseCardReducer,
        courseDetail: courseDetailReducer,
        header: headerReducer,
        footer: footerReducer,
        pdf: pdfReducer,
    },
})