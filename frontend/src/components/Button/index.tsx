import React, { useCallback } from 'react'

interface Props {
    buttonContent?: string
    onClick?: () => void
}

const Button: React.FunctionComponent<Props> = ({ buttonContent, onClick = () => ({}) }) => {
    const handleClick = useCallback(() => {
        onClick()
    }, [onClick])

    return (
        <>
          <button onClick={handleClick}>{buttonContent}</button>
        </>
    );
};

export default Button;