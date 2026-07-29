import "./style.css"
import {useNavigate} from "react-router";

const PageNotFound = () => {
    const navigate = useNavigate()
    return (
        <div className="custom-bg text-dark">
            <div className="errors-page-wrapper">
                <div className="errors-page">
                    <h1 className="display-1 fw-bold">404</h1>
                    <p className="fs-2 fw-medium mt-4">Страница не найдена.</p>
                    <p className="mt-4 mb-5">Страница, которую вы ищете, отсутствует или переехала.</p>
                    <button onClick={()=>navigate("/")}>
                        Вернуться на главную
                    </button>
                </div>
            </div>
        </div>
    )
}

export default PageNotFound;