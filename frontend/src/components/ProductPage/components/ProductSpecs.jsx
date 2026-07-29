import {memo, useState} from "react";
import "./styles.css"
import {config} from "../../../config";


const ProductSpecs = ({itemCard}) =>{
    const item = itemCard
    const [activeTab, setActiveTab] = useState('description')
    const [props, setProps] = useState([])
    const [inputFields, setInputFields] = useState([])

    const handleAddInput = (id, propName) => {
       const existingIndex = inputFields.findIndex(field => field.id === id);
       if (existingIndex !== -1) {
           const updated = [...inputFields]
           updated.splice(existingIndex, 1)
           setInputFields(updated)
       } else {
           setInputFields([...inputFields, {id, name:propName, value: ''}])
       }
    }

    const handleInputChange = (index, event) => {
        const values = [...inputFields];
        values[index].value = event.target.value;
        setInputFields(values);
    };

    const togglePropVisible = () =>{
        setActiveTab('prop')
    }
    const toggleDescriptionVisible = () =>{
        setActiveTab('description')
    }

    const toggleSpecVisible = () =>{
        setActiveTab('spec')
    }

    const handleGetProps = async() => {
        try {
            const response = await fetch (`${config.apiUrl}/admin/props`, {
               method : "GET",
                headers :{
                    "Content-Type": "application/json"
                }
            })

            if (response.ok) {
                const result = await response.json()
                setProps(result)
            }
        } catch (error) {
            throw error
        }
    }

    const handleUploadSpecs = async() => {
        const payload = inputFields.map(field => ({
            prop_id : field.id,
            prop_value : field.value
        }))
        try {
            const response = await fetch (`${config.apiUrl}/admin/newprops/${itemCard.itemData.id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(payload)
            })


        } catch (err) {
            throw err
        }


    }




    return(
        <div className="product-info-spec-wrapper">
            <div className ="product-info-spec__btn-section">
                <button type = "button"
                        onClick = {toggleDescriptionVisible}
                        className={activeTab === 'description' ? 'active' : ''}
                >Описание</button>
                <button type = "button"
                        onClick = {togglePropVisible}
                        className={activeTab === 'prop' ? 'active' : ''}
                >Характеристики</button>
                <button type = "button"
                        onClick={toggleSpecVisible}
                        className={activeTab === 'spec' ? 'active' : ''}
                >Документация</button>
            </div>
            {activeTab === 'description' && (
                <div>
                    <h1>{item.itemData.name}</h1>
                    <p>{item.itemData.item_description}</p>
                </div>
            )}

            {activeTab === 'prop' && (
                <div className="product-info-props">
                    {item.itemData.prop_name_array ? (
                    item.itemData.prop_name_array.map((propName, index) => (
                        <div className="product-info-row" key={index}>
                            <div className="product-info-name">
                                {propName}
                            </div>
                            <div className="product-info-value">
                                {item.itemData.prop_value_array[index]}
                            </div>
                        </div>
                    ))
                        ) : (
                            <div className = "product-specs-add-admin-wrapper">
                                {props.length === 0 && (
                                    <>
                            <h1>Характеристики отсутствуют</h1>

                                    </>
                                    )}
                        <button onClick = {handleGetProps} id = {"add-btn"}>Добавить характеристики</button>
                                <div className= "product-specs-add-buttons-wrapper">
                                {props.length > 0 && props.map((item, index) => (
                                    item.trim() !== '' && (
                                       <button className = {inputFields.some(f=> f.id ===index) ? "active" : ""} onClick = {() => handleAddInput(index,item)}>{item}</button>
                                    ))
                                    )}
                                    </div>
                                {props.length >0  && (
                                    <>
                                        <h3>Все характеристики вводятся строго через запятую!</h3>

                                    </>
                                )}
                                {inputFields.map((field, value) => (
                                    <div className="product-spec-add-inputs-wrapper">
                                    <div className ="product-spec-add-inputs">
                                    <label>{field.name}: </label>
                                    <input type = "text"
                                           value = {field.value}
                                           onChange={(event)=>handleInputChange(value,event)}
                                           placeholder="Введите характеристики" />
                                    </div>
                                        </div>
                                ))}
                                {inputFields.length >0 && (

                                <button onClick = {handleUploadSpecs} id = "send-btn">Отправить</button>
                                    )}
                            </div>

                            )}
                </div>
            )}
            {activeTab === 'spec' &&  (
                <div className = "product-files-wrapper">
                    {item.specData?.map((spec) => (
                        <div className = "product-files" key ={spec.spec_file_name}>
                            <img alt = "icon" src = {`/${spec.spec_file_pic}`}></img>
                            <a href = {`${config.docServerUrl}${spec.spec_file_link}`} target="_blank" rel ="noreferrer">{spec.spec_file_name} {item.itemData.name}</a>
                        </div>
                    ))}

                </div>
            )}

        </div>
    )
}

export default memo(ProductSpecs);