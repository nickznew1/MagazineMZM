import {useState} from "react";
import {useForm} from "react-hook-form";
import {config} from "../../../config";


const AdminPageFilterCreateItem = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm({
        mode: "onBlur",
        values:{
            name:"",
            price:"",
            item_type:"",
            secondary_type:"",
            item_description:"",
            item_short_description:"",
            article:"",
            document_name: "",
        }
    });
    const [previewImg, setPreviewImg] = useState(null);
    const [previewPdf, setPreviewPdf] = useState(null);
    const [img, setImg] = useState(null);
    const [pdf, setPdf] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(false);
    const [message, setMessage] = useState(null);

    const handleImageChange = (event) =>{

        const image = event.target.files[0]

        setImg(image)
        if (!image) return;

        const reader = new FileReader();
        reader.onload = (event) => {
            setPreviewImg(event.target.result);
        }
        reader.readAsDataURL(image);

    }

    const  handlePdfChange = (event) =>{
        const pdf = event.target.files[0];
        setPdf(pdf);
        if (!pdf) return;

        const reader = new FileReader();
        reader.onload = (event) => {
            setPreviewPdf(event.target.result);
        }
        reader.readAsDataURL(pdf)
    }

    const handleSubmitNewItem = async (data) => {
        setLoading(true);
        setError(false);
        const formData = new FormData();

        Object.keys(data).forEach((key) => {
            formData.append(key, data[key])
        })
        if (previewImg){
            formData.append('imgFile',img)
        }
        if (previewPdf){

            formData.append('pdfFile', pdf);
        }



        try {
            const response = await fetch(`${config.apiUrl}/item/create`, {
                method: "POST",
                body: formData,
            })
            if (response.ok) {

                setMessage(true)
            } else {
                setLoading(false);
                setError(true);
                setMessage(false)
            }
        } catch (error) {
            throw new Error (error)
        }

    }





    return (
        <form onSubmit={handleSubmit(handleSubmitNewItem)} >
        <table className="admin_page-item-create-table">
            <thead>
            <tr>
                <th>Загрузить картинку</th>

                <td>
                    <label form="userfile">Выберите файл:</label>
                    <input
                        type="file"
                        onChange={handleImageChange}
                        id = "userfile"
                        required='Обязательное поле'
                    />
                </td>

                {previewImg &&
                    <td>
                        <img src={previewImg} alt="" />
                    </td>
                }
            </tr>
            <tr>
                <th>Название</th>
                <td> <input type="text"
                             placeholder ="Введите название"
                             style={{ borderColor: errors.name ? 'red' : '#ccc'  }}
                             {...register("name", {
                                 required: 'Обязательное поле',
                             })
                             }>
                </input>
                    {errors.name && <span className="admin-page-errors">{errors.name.message}</span> }
                </td>
            </tr>
            <tr>
                <th>Тип/Способ измерения</th>
                <td><input type="text"
                           placeholder ="Введите тип"
                           style={{ borderColor: errors.type ? 'red' : '#ccc'  }}
                           {...register("item_type", {
                               required: 'Обязательное поле',
                           })
                           }>
                </input>
                    {errors.type && <span className="admin-page-errors">{errors.type.message}</span> }
                    <div>
                        <h3>Существующие типы:</h3>
                        <p>Датчик давления</p>
                        <p>Датчик расхода</p>
                        <p>Датчик уровня</p>
                    </div>
                </td>
                <td><input type="text"
                           placeholder ="Введите способ"
                           style={{ borderColor: errors.secondaryType ? 'red' : '#ccc'  }}
                           {...register("secondary_type", {
                               required: 'Обязательное поле',
                           })
                           }>
                </input>
                    {errors.secondaryType && <span className="admin-page-errors">{errors.secondaryType.message}</span> }
                </td>
            </tr>
            <tr>
                <th>Описание</th>
                <td><textarea
                           placeholder ="Введите описание"
                           style={{ borderColor: errors.description ? 'red' : '#ccc'  }}
                           {...register("item_description", {
                               required: 'Обязательное поле',
                           })
                           }>
                </textarea>
                    {errors.description && <span className="admin-page-errors">{errors.description.message}</span> }
                </td>
            </tr>
            <tr>
                <th>Краткое описание</th>
                <td><textarea
                           placeholder ="Введите краткое описание"
                           style={{ borderColor: errors.shortDescription ? 'red' : '#ccc'  }}
                           {...register("item_short_description", {
                               required: 'Обязательное поле',
                           })
                           }>
                </textarea>
                    {errors.shortDescription && <span className="admin-page-errors">{errors.shortDescription.message}</span> }
                </td>
            </tr>
            <tr>
                <th>Начальная цена</th>
                <td><input type="text"
                           placeholder ="Введите начальную цену"
                           style={{ borderColor: errors.price ? 'red' : '#ccc'  }}
                           {...register("price", {
                               required: 'Обязательное поле',
                           })
                           }>
                </input>
                    {errors.price && <span className="admin-page-errors">{errors.price.message}</span> }
                </td>
            </tr>
            <tr>
                <th>Артикуль</th>
                <td> <input type= "text"
                            placeholder ="Введите артикуль"
                            style={{ borderColor: errors.article ? 'red' : '#ccc'  }}
                            {...register("article", {
                                required: 'Обязательное поле',
                            })
                            }>
                </input>
                    {errors.article && <span className="admin-page-errors">{errors.article.message}</span> }
                </td>
            </tr>
            <tr>
                <th>Документы</th>
                <td>
                <label form="userfile">Выберите файл:</label>
                    <input
                        type="file"
                        onChange={handlePdfChange}
                        id = "userfile"
                        required='Обязательное поле'
                    />
                    <input type="text"
                           placeholder ="Название файла"
                           style={{ borderColor: errors.file ? 'red' : '#ccc'  }}
                           {...register("document_name", {
                               required: 'Обязательное поле',
                           })
                           }>
                    </input>
                    {errors.file && <span className="admin-page-errors">{errors.file.message}</span> }
                </td>
            </tr>
            </thead>
        </table>
            <p>*Товар изначально не открывается обычным пользователям. Чтобы включить отображение, необходимо перейти на карточку товара и выбрать соответствующий пункт</p>
            <p>**Характеристики прописываются там же.(обязательно)</p>
            <button type = "submit"> Отправить</button>
            {message ? (
                <p>Товар успешно загружен!</p>
            ) : (
                <p>Ошибка при отправлении товара</p>
            )}

        </form>
    )
}

export default AdminPageFilterCreateItem