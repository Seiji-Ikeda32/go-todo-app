interface Props {
    titleName?: string
}

const Title: React.FunctionComponent<Props> = ({ titleName }) => {
    return (
        <>
          <h1>{titleName}</h1>
        </>
    );
};

export default Title;