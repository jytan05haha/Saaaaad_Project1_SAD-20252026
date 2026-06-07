import { Container, Title, Button } from '@mantine/core'
import { Link } from 'react-router-dom'

export default function LoggedInFlow(){
  function logout(){
    localStorage.removeItem('token');
    location.reload();
  }

  return (
    <Container style={{paddingTop:24}}>
      <Title order={2}>Dashboard</Title>
      <Button component={Link} to="/orders" style={{marginRight:8}}>Orders</Button>
      <Button color="red" onClick={logout}>Sign out</Button>
    </Container>
  )
}
