import { useAccessors } from '../../../store'
import { useEffect } from 'react'
import './Accessors.css'

const Accessors = () => {
    const accessors = useAccessors(state => state.accessors)
    const data = useAccessors(state => state.readAccessors)
    useEffect(() => {
        data()
    }, [])

    return <>
        <table>
            <thead>
                <tr key={42}>
                    <th>Код</th>
                    <th>Идентификатор каталога</th>
                    <th>Идентификатор номеклатуры</th>
                    <th>Наименование</th>
                    <th>Идентификатор модификации</th>
                    <th>Идентификатор</th>
                    <th>Версия</th>
                    <th>Действия</th>
                </tr>
            </thead>
            <tbody>
                {
                    accessors.map((accessor, row) => {
                        return <tr key={ row }>
                            <td>{ accessor.rn }</td>
                            <td>{ accessor.crn }</td>
                            <td>{ accessor.version }</td>
                            <td>{ accessor.code }</td>
                            <td>{ accessor.name }</td>
                            <td>{ accessor.dicnomns }</td>
                            <td>{ accessor.nommodif }</td>
                            <td>
                                <button>+</button>
                                <button>-</button>
                            </td>
                        </tr>
                    })
                }
            </tbody>
        </table>
    </>
}

export default Accessors