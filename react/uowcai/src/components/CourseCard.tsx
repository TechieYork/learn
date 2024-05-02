import { Card, Skeleton, Image } from "antd";
import { useSelector, useDispatch } from "react-redux";
import { setImageLoaded } from "./courseCardSlice";
import { UOWCAI, UOWCAI_COURSE_DEFAULT_IMAGE } from "../constants/Course";
const { Meta } = Card;

interface CourseInfo {
  title: string;
  description?: string;
  image?: string;
  url: string;
  onClick?: () => void;
}

function CourseCard(course: CourseInfo) {
  let imageLoaded = useSelector((state: any) => state.courseCard.imageLoaded);
  let dispatch = useDispatch();

  return (
    <Card
      hoverable
      style={{ minWidth: "300px", minHeight: "300px" }}
      cover={
        <Image
          alt={course.title}
          src={course.image || UOWCAI_COURSE_DEFAULT_IMAGE}
          style={{
            opacity: imageLoaded ? 1 : 0,
            transition: "opacity 0.3s ease-in-out",
            backgroundColor: "lightgray",
            height: "300px",
            width: "100%",
            border: "1px solid #f0f0f0",
            objectFit: "cover",
          }}
          preview={false}
          onLoad={() => dispatch(setImageLoaded(true))}
          fallback={UOWCAI_COURSE_DEFAULT_IMAGE}
          onClick={course.onClick}
        />
      }
    >
      <Meta title={course.title} description={course.description} />
    </Card>
  );
}

export { CourseCard };
