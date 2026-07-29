import "./style.css"
import {useNavigate} from "react-router";

const ForbiddenPage = () => {
    const navigate = useNavigate()
    return (
        <>
            <div className="custom-bg text-dark">
                <div className="errors-page-wrapper">
                    <div className="errors-page">
                        <h1 className="display-1 fw-bold">403</h1>
                        <p className="fs-2 fw-medium mt-4">У вас нет прав на посещение данной страницы.</p>
                        <p className="mt-4 mb-5">Вы не имеете прав администратора.</p>
                        <button onClick={() => navigate("/")}>
                            Вернуться на главную
                        </button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default ForbiddenPage;