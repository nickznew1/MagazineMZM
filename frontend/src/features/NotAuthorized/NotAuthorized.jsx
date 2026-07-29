import "./style.css"
import {useNavigate} from "react-router";

const ForbiddenPage = () => {
    const navigate = useNavigate()
    return (
        <>
            <div className="custom-bg text-dark">
                <div className="errors-page-wrapper">
                    <div className="errors-page">
                        <h1 className="display-1 fw-bold">401</h1>
                        <p className="fs-2 fw-medium mt-4">Вы не авторизованы.</p>
                        <p className="mt-4 mb-5">Вы должны авторизоваться для доступа к этой странице.</p>
                        <button onClick={() => navigate("/auth")}>
                            Вернуться на главную
                        </button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default ForbiddenPage;