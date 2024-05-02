import ListGroup from "./components/ListGroup";
import Alert from "./components/Alert";

function App() {
  let items = ["item1", "item2", "item3"];

  const handleSelectItem = (item: string) => {
    console.log(item);
  };

  return (
    <div>
      <ListGroup
        items={items}
        heading="List Heading..."
        onSelectItem={handleSelectItem}
      />
      <Alert>Alert OPS~~~~</Alert>
    </div>
  );
}

export default App;
