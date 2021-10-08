import { useRef, useState } from "react";
import type { NextPage } from "next";
import Link from "next/link";
import { useRouter } from "next/router";
import {
  Button,
  Card,
  Col,
  Form,
  InputGroup,
  ListGroup,
  Row,
  Table,
} from "react-bootstrap";
import { FaPaperPlane, FaPlus, FaTrash, FaTint } from "react-icons/fa";

import ButtonIcon from "../../components/ButtonIcon";
import Layout from "../../components/Layout";
import ModalContainer from "../../components/ModalContainer";
import Panel from "../../components/Panel";
import TableTaskItem from "../../components/TableTaskItem";
import useModal from "../../hooks/useModal";
import { cropsData, tasksData, notesData } from "../../data";

const ReservoirDetail: NextPage = () => {
  const router = useRouter();
  const { id } = router.query;
  const { modalOpen, showModal, closeModal } = useModal();
  const [dueDate, setDueDate] = useState("");
  const [priority, setPriority] = useState("");
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [selectedCategory, setSelectedCategory] = useState();
  const [isError, setIsError] = useState(false);
  const target = useRef(null);

  const addTaskReservoir = () => {
    if (!dueDate || !priority || !title) {
      setIsError(true);
    } else {
      setIsError(false);
      closeModal();
    }
  };

  return (
    <Layout>
      <Row>
        <Col>
          <h3 className="pb-3">
            Органический салат
            <small className="ms-2 text-muted">Рост</small>
          </h3>
        </Col>
      </Row>
      <Row>
        <Col className="mb-3">
          <ButtonIcon
            label="Add Task"
            icon={<FaPlus className="me-2" />}
            onClick={showModal}
            variant="primary"
          />
          <ButtonIcon
            label="Watering"
            icon={<FaTint className="me-2" />}
            onClick={() => {}}
            variant="secondary"
            textColor="text-light"
          />
        </Col>
      </Row>
      <Row>
        <Col md={5} className="mb-3">
          <Card>
            <Card.Img variant="top" src="/no-img.png" />
            <Card.Body>
              <Card.Title>Basic Info</Card.Title>
              <ListGroup>
                <ListGroup.Item>
                  <span className="text-muted">Площадь Га</span>
                  <span className="ms-3">1</span>
                </ListGroup.Item>
                <ListGroup.Item>
                  <span className="text-muted">Локация</span>
                  <span className="ms-3">Поле (Outdoor)</span>
                </ListGroup.Item>
                <ListGroup.Item>
                  <span className="text-muted">Партия</span>
                  <span className="ms-3">2</span>
                </ListGroup.Item>
                <ListGroup.Item>
                  <span className="text-muted">Сбор урожая</span>
                  <span className="ms-3">1</span>
                </ListGroup.Item>
                <ListGroup.Item>
                  <span className="text-muted">Источник воды</span>
                  <span className="ms-3">Река</span>
                </ListGroup.Item>
              </ListGroup>
            </Card.Body>
          </Card>
        </Col>
        <Panel title="Crops" md={7} lg={7}>
          <Table responsive>
            <thead>
              <tr>
                <th>Сбор урожая</th>
                <th>Идентификатор партии </th>
                <th>Дата посадки</th>
                <th>Дней с момента посева</th>
                <th>Количество </th>
                <th>Последний полив</th>
              </tr>
            </thead>
            <tbody>
              {cropsData &&
                cropsData.map(
                  ({
                    id,
                    varieties,
                    batchId,
                    seedingDate,
                    daysSinceSeeding,
                    qty,
                    qtyUnit,
                    lastWatering,
                  }) => (
                    <tr key={id}>
                      <td>
                        <Link href={`/crops/${id}`}>{varieties}</Link>
                      </td>
                      <td>{batchId}</td>
                      <td>{seedingDate}</td>
                      <td>{daysSinceSeeding}</td>
                      <td>{`${qty} ${qtyUnit}`}</td>
                      <td>{lastWatering}</td>
                    </tr>
                  )
                )}
            </tbody>
          </Table>
        </Panel>
      </Row>
      <Row>
        <Panel title="Notes" md={6} lg={6}>
          <>
            <InputGroup className="mb-3">
              <Form.Control type="text" placeholder="Создайте заметку" />
              <Button variant="secondary">
                <div className="d-flex align-items-center">
                  <FaPaperPlane />
                </div>
              </Button>
            </InputGroup>
            <ListGroup>
              {notesData &&
                notesData.map(({ id, title, createdOn }) => (
                  <ListGroup.Item key={id}>
                    <div className="d-flex align-items-center justify-content-between py-1">
                      <div>
                        <div className="mb-1">{title}</div>
                        <small className="text-muted">{createdOn}</small>
                      </div>
                      <div>
                        <FaTrash />
                      </div>
                    </div>
                  </ListGroup.Item>
                ))}
            </ListGroup>
          </>
        </Panel>
        <Panel title="Tasks" md={6} lg={6}>
          <Table responsive>
            <thead>
              <tr>
                <th className="w-75">Наименование</th>
                <th>Category</th>
              </tr>
            </thead>
            <tbody>
              {tasksData &&
                tasksData.map(
                  ({ id, item, details, dueDate, priority, category }) => (
                    <tr key={id}>
                      <td>
                        <TableTaskItem
                          id={id}
                          item={item}
                          details={details}
                          dueDate={dueDate}
                          priority={priority}
                        />
                      </td>
                      <td>
                        <span className="text-uppercase">{category}</span>
                      </td>
                    </tr>
                  )
                )}
            </tbody>
          </Table>
        </Panel>
      </Row>
      <ModalContainer
        title="Область: Добавить новую задачу по органическому салату"
        isShow={modalOpen}
        handleCloseModal={closeModal}
        handleSubmitModal={addTaskReservoir}
      >
        <>
          <Form>
            <Form.Group className="mb-3">
              <Form.Label>Due Date</Form.Label>
              <InputGroup ref={target}>
                <Form.Control
                  type="date"
                  value={dueDate}
                  onChange={(e) => setDueDate(e.target.value)}
                />
              </InputGroup>
              {isError && (
                <Form.Text className="text-danger">
                  The due date field is required
                </Form.Text>
              )}
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>Is this task urgent?</Form.Label>
              <Form.Check
                type="radio"
                label="Yes"
                name="priority"
                onChange={() => setPriority("urgent")}
              />
              <Form.Check
                type="radio"
                label="No"
                name="priority"
                onChange={() => setPriority("normal")}
              />
              {isError && (
                <Form.Text className="text-danger">
                  The priority field is required
                </Form.Text>
              )}
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>Task Category</Form.Label>
              <Form.Select
                onChange={(e) => setSelectedCategory(e.target.value)}
              >
                <option>Please select category</option>
                <option value="1">Reservoir</option>
                <option value="2">Pest Control</option>
                <option value="3">Safety</option>
                <option value="4">Sanitation</option>
              </Form.Select>
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>Title</Form.Label>
              <Form.Control
                type="text"
                onChange={(e) => setTitle(e.target.value)}
                value={title}
              />
              {isError && (
                <Form.Text className="text-danger">
                  The title field is required
                </Form.Text>
              )}
            </Form.Group>
            <Form.Group className="mb-3">
              <Form.Label>Description</Form.Label>
              <Form.Control
                as="textarea"
                onChange={(e) => setDesc(e.target.value)}
                value={desc}
                style={{ height: "120px" }}
              />
            </Form.Group>
          </Form>
        </>
      </ModalContainer>
    </Layout>
  );
};

export default ReservoirDetail;
