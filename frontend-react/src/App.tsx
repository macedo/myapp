import {
  CapitalOsAuthenticationProvider,
  CardsApp,
  AccountDetails,
} from "@capitalos/react";

function App() {
  const getToken = async () => {

    const response = await fetch("/api/get-capitalos-token")
    const data = await response.json()
    console.log("data", data)
    return data.token
  }

  return (
    <CapitalOsAuthenticationProvider getToken={getToken}>
      <AppContent />
    </CapitalOsAuthenticationProvider>
  )
}

function AppContent() {
  return (
    <div>
      <div className="page">
        <CardsApp />
      </div>
    </div>
  )
}

export default App
