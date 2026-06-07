import { Container, Title, Grid, Card, Button } from '@mantine/core'
import { Link } from 'react-router-dom'

export default function Home(){
  return (
    <Container style={{paddingTop:24}}>
      <Title order={2}>Welcome to SaLaad Food</Title>
      <Grid>
        <Grid.Col span={6}>
          <Card shadow="sm">
            <Title order={4}>Restaurants</Title>
            <p>Browse restaurants and place orders.</p>
            <Button component={Link} to="/orders">Browse Orders</Button>
          </Card>
        </Grid.Col>
      </Grid>
    </Container>
  )
}
