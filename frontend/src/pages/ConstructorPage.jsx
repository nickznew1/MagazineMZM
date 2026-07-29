import ItemConstructor from "../components/ConstructorPage/components/ItemConstructor";
import {GetItemConstructor} from "../components/ConstructorPage/hooks/getItemConstructor";
import {useParams} from "react-router";
import PageNotFound from "../features/404NotFound/PageNotFound";
import NotAuthorized from "../features/NotAuthorized/NotAuthorized";
import LoadingScreen from "../features/LoadingScreen/LoadingScreen";


const ConstructorPage = () => {
    const {id} = useParams()
    const constructor = GetItemConstructor(id)

    if (constructor.loading){
        return <div>loading...</div>

    }
    if (constructor.error){
        return(
            <NotAuthorized/>
        )
    }

    return (
        <>
            <ItemConstructor
               constructor = {constructor}
            />
        </>
    )
}

export default ConstructorPage