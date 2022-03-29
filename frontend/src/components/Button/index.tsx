import React, { useCallback } from 'react'

interface Props {
    onClick?: () => void
}

const Button: React.FunctionComponent<Props> = ({ children, onClick = () => ({}) }) => {
    const handleClick = useCallback(() => {
        onClick()
    }, [onClick])

    return (
        <>
          <button onClick={handleClick}>{children}</button>
        </>
    );
};

export default Button;