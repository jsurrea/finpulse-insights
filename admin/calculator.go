package main

import (
	"fmt"
	"math"
	"strings"
)

var BrokerageCredibility = map[string]float64{
    // Tier 1: Bulge Bracket Banks (0.90-0.95)
    "The Goldman Sachs Group":           0.95,
    "JPMorgan Chase & Co.":             0.95,
    "Morgan Stanley":                   0.94,
    "Bank of America":                  0.93,
    "Citigroup":                        0.92,
    "Credit Suisse Group":              0.90,
    "Deutsche Bank Aktiengesellschaft": 0.90,
    
    // Tier 2: Major Investment Banks (0.85-0.89)
    "Wells Fargo & Company":            0.88,
    "Barclays":                         0.87,
    "UBS Group":                        0.87,
    "HSBC":                            0.86,
    "BNP Paribas":                     0.85,
    "Macquarie":                       0.85,
    
    // Tier 3: Regional/Specialized Major Banks (0.80-0.84)
    "Royal Bank of Canada":            0.84,
    "Royal Bank Of Canada":            0.84, // Variante
    "BMO Capital Markets":             0.83,
    "Scotiabank":                      0.83,
    "TD Securities":                   0.82,
    "CIBC":                           0.82,
    "Mizuho":                         0.81,
    "Daiwa Capital Markets":          0.80,
    
    // Tier 4: Established Investment Banks (0.75-0.79)
    "Jefferies Financial Group":       0.79,
    "Evercore ISI":                    0.78,
    "Guggenheim":                      0.78,
    "Oppenheimer":                     0.77,
    "Raymond James":                   0.77,
    "Raymond James Financial":         0.77,
    "Piper Sandler":                   0.76,
    "Truist Financial":                0.76,
    "KeyCorp":                         0.75,
    "Stifel Nicolaus":                 0.75,
    
    // Tier 5: Mid-Tier Investment Banks (0.70-0.74)
    "Wedbush":                         0.74,
    "Robert W. Baird":                 0.73,
    "Canaccord Genuity Group":         0.72,
    "TD Cowen":                        0.72,
    "Needham & Company LLC":           0.71,
    "Craig Hallum":                    0.71,
    "Janney Montgomery Scott":         0.70,
    "Stephens":                        0.70,
    
    // Tier 6: Specialized/Boutique Banks (0.65-0.69)
    "HC Wainwright":                   0.69,
    "B. Riley":                        0.68,
    "Rosenblatt Securities":           0.68,
    "JMP Securities":                  0.67,
    "BTIG Research":                   0.67,
    "Sanford C. Bernstein":            0.67,
    "Leerink Partners":                0.66,
    "Cantor Fitzgerald":               0.66,
    "Roth Capital":                    0.65,
    "Roth Mkm":                        0.65,
    
    // Tier 7: Smaller Specialized Firms (0.60-0.64)
    "Telsey Advisory Group":           0.64,
    "Wolfe Research":                  0.63,
    "Susquehanna":                     0.63,
    "Keefe, Bruyette & Woods":         0.62,
    "DA Davidson":                     0.62,
    "Jones Trading":                   0.61,
    "Benchmark":                       0.61,
    "Compass Point":                   0.60,
    "Sidoti":                          0.60,
    
    // Tier 8: Research-Focused/Niche Firms (0.55-0.59)
    "Maxim Group":                     0.59,
    "Noble Financial":                 0.58,
    "Tigress Financial":               0.58,
    "Alliance Global Partners":        0.57,
    "Northland Securities":            0.57,
    "Longbow Research":                0.56,
    "Argus":                           0.56,
    "Lake Street Capital":             0.55,
    "Vertical Research":               0.55,
    
    // Tier 9: Smaller/Regional Firms (0.50-0.54)
    "Chardan Capital":                 0.54,
    "Ascendiant Capital Markets":      0.53,
    "BWS Financial":                   0.53,
    "ThinkEquity":                     0.52,
    "Westpark Capital":                0.52,
    "Barrington Research":             0.51,
    "Capital One Financial":           0.51,
    "Monness Crespi & Hardt":          0.50,
    
    // Tier 10: Emerging/Specialized Firms (0.45-0.49)
    "Hovde Group":                     0.49,
    "D. Boral Capital":                0.48,
    "Loop Capital":                    0.48,
    "Redburn Atlantic":                0.47,
    "BNP Paribas Exane":              0.47,
    "Berenberg Bank":                  0.46,
    "Arete Research":                  0.46,
    "Arete":                           0.46,
    "Citizens Jmp":                    0.45,
    "Glj Research":                    0.45,
    "LADENBURG THALM/SH SH":          0.45,
    "Melius":                          0.45,
    "Zelman & Associates":             0.45,
    "CICC Research":                   0.45,
    "Fundamental Research":            0.45,
    "DNB Markets":                     0.45,
    "DZ Bank":                         0.45,
    "Cfra":                            0.45,
    "Greenridge Global":               0.45,
    "Siebert Williams Shank":          0.45,
    "Northcoast Research":             0.45,
}


var RatingScore = map[string]int{
    // Empty/Unknown
    "": 3,
    
    // Strong Buy Ratings (5)
    "Strong-Buy":      5,
    "Top Pick":        5,
    "Action List Buy": 5,
    
    // Buy Ratings (4)
    "Buy":             4,
    "Speculative Buy": 4,
    "Moderate Buy":    4,
    "Outperform":      4,
    "Outperformer":    4,
    "Market Outperform": 4,
    "Mkt Outperform":  4,
    "Sector Outperform": 4,
    "Overweight":      4,
    "Positive":        4,
    
    // Neutral/Hold Ratings (3)
    "Hold":            3,
    "Neutral":         3,
    "Equal Weight":    3,
    "Market Perform":  3,
    "Sector Perform":  3,
    "Peer Perform":    3,
    "In-Line":         3,
    "Sector Weight":   3,
    
    // Underperform/Sell Ratings (1-2)
    "Underperform":    2,
    "Sector Underperform": 2,
    "Underweight":     2,
    "Reduce":          2,
    "Sell":            1,
    "Strong Sell":     1,
}

// getActionScore analiza el tipo de acción y asigna un score
func getActionScore(action string) float64 {
    switch strings.ToLower(action) {
    case "upgraded by":
        return 0.9 // Muy positivo
    case "initiated by":
        return 0.8 // Positivo (nueva cobertura)
    case "target raised by":
        return 0.7 // Positivo
    case "target set by":
        return 0.6 // Neutral-positivo
    case "reiterated by":
        return 0.5 // Neutral
    case "target lowered by":
        return 0.3 // Negativo
    case "downgraded by":
        return 0.1 // Muy negativo
    default:
        return 0.5 // Default neutral
    }
}


// calculateRecommendation determina BUY/SELL/HOLD basado en múltiples factores
func calculateRecommendation(rec *StockRecommendation) string {
    // Normalizar ratings vacíos
    ratingFrom := rec.RatingFrom
    ratingTo := rec.RatingTo
    if ratingFrom == "" {
        ratingFrom = "Neutral" // Asumir neutral si está vacío
    }
    
    ratingFromScore := RatingScore[ratingFrom]
    ratingToScore := RatingScore[ratingTo]
    
    // 1. Rating final (peso: 40%)
    finalRatingWeight := 0.4
    var finalRatingRec string
    if ratingToScore >= 4 {
        finalRatingRec = "BUY"
    } else if ratingToScore <= 2 {
        finalRatingRec = "SELL"
    } else {
        finalRatingRec = "HOLD"
    }
    
    // 2. Cambio de rating (peso: 30%)
    ratingChange := ratingToScore - ratingFromScore
    ratingChangeWeight := 0.3
    var ratingChangeRec string
    if ratingChange > 0 {
        ratingChangeRec = "BUY"
    } else if ratingChange < 0 {
        ratingChangeRec = "SELL"
    } else {
        ratingChangeRec = "HOLD"
    }
    
    // 3. Análisis de acción (peso: 20%)
    actionScore := getActionScore(rec.Action)
    actionWeight := 0.2
    var actionRec string
    if actionScore >= 0.7 {
        actionRec = "BUY"
    } else if actionScore <= 0.3 {
        actionRec = "SELL"
    } else {
        actionRec = "HOLD"
    }
    
    // 4. Target price analysis (peso: 10%)
    targetWeight := 0.1
    var targetRec string = "HOLD"
    if rec.TargetTo > 0 && rec.TargetFrom > 0 {
        targetChange := (rec.TargetTo - rec.TargetFrom) / rec.TargetFrom
        if targetChange > 0.05 { // 5% increase
            targetRec = "BUY"
        } else if targetChange < -0.05 { // 5% decrease
            targetRec = "SELL"
        }
    }
    
    // Calcular score ponderado
    scores := map[string]float64{
        "BUY":  0.0,
        "SELL": 0.0,
        "HOLD": 0.0,
    }
    
    scores[finalRatingRec] += finalRatingWeight
    scores[ratingChangeRec] += ratingChangeWeight
    scores[actionRec] += actionWeight
    scores[targetRec] += targetWeight
    
    // Encontrar la recomendación con mayor score
    maxScore := 0.0
    result := "HOLD"
    for recommendation, score := range scores {
        if score > maxScore {
            maxScore = score
            result = recommendation
        }
    }
    
    return result
}


// calculateConfidence calcula la confianza basada en múltiples factores
func calculateConfidence(rec *StockRecommendation) float64 {
    // 1. Credibilidad del brokerage (35%)
    brokerageScore := BrokerageCredibility[rec.Brokerage]
    if brokerageScore == 0 {
        brokerageScore = 0.45 // Default para brokerages no listados
    }
    
    // 2. Consistencia de rating (25%)
    ratingFromScore := RatingScore[rec.RatingFrom]
    ratingToScore := RatingScore[rec.RatingTo]
    if ratingFromScore == 0 {
        ratingFromScore = 3 // Neutral default
    }
    
    consistencyScore := 0.5
    ratingChange := math.Abs(float64(ratingToScore - ratingFromScore))
    
    if ratingChange >= 2 {
        consistencyScore = 0.9 // Cambio significativo = alta confianza
    } else if ratingChange == 1 {
        consistencyScore = 0.7
    } else if ratingChange == 0 {
        // Rating mantenido - depende de la acción
        if strings.Contains(rec.Action, "reiterated") {
            consistencyScore = 0.8 // Reiterado = confianza alta
        }
    }
    
    // 3. Claridad de la acción (25%)
    actionScore := getActionScore(rec.Action)
    actionClarityScore := actionScore
    if actionScore == 0.5 { // Neutral actions
        actionClarityScore = 0.6
    }
    
    // 4. Target price reliability (15%)
    targetScore := 0.5
    if rec.TargetTo > 0 && rec.TargetFrom > 0 {
        pctChange := math.Abs((rec.TargetTo - rec.TargetFrom) / rec.TargetFrom)
        if pctChange > 0.1 {
            targetScore = 0.8
        } else if pctChange > 0.05 {
            targetScore = 0.7
        } else if pctChange == 0 {
            targetScore = 0.6
        }
    } else if rec.TargetTo > 0 {
        targetScore = 0.7 // At least has a target
    }
    
    // Calcular confianza final ponderada
    confidence := (brokerageScore * 0.35) + 
                 (consistencyScore * 0.25) + 
                 (actionClarityScore * 0.25) + 
                 (targetScore * 0.15)
    
    return math.Max(0.1, math.Min(0.99, confidence))
}

// calculateScore genera un score numérico para ranking
func calculateScore(rec *StockRecommendation) int {
    score := 50 // Base score
    
    // Rating final
    finalRatingScore := RatingScore[rec.RatingTo]
    score += (finalRatingScore - 3) * 20 // -40 a +40
    
    // Cambio de rating
    ratingChange := RatingScore[rec.RatingTo] - RatingScore[rec.RatingFrom]
    score += ratingChange * 15 // -60 a +60
    
    // Target upside (si disponible)
    if rec.TargetTo > 0 && rec.TargetFrom > 0 {
        pctChange := (rec.TargetTo - rec.TargetFrom) / rec.TargetFrom
        score += int(pctChange * 100) // Porcentaje como puntos
    }
    
    // Credibilidad del brokerage
    brokerageMultiplier := BrokerageCredibility[rec.Brokerage]
    if brokerageMultiplier == 0 {
        brokerageMultiplier = 0.5
    }
    score = int(float64(score) * brokerageMultiplier)
    
    return score
}

// generateReason crea una explicación textual
func generateReason(rec *StockRecommendation) string {
    parts := []string{}
    
    // Acción principal
    if strings.Contains(strings.ToLower(rec.Action), "upgrade") {
        parts = append(parts, fmt.Sprintf("%s upgraded %s from %s to %s", 
            rec.Brokerage, rec.Ticker, rec.RatingFrom, rec.RatingTo))
    } else if strings.Contains(strings.ToLower(rec.Action), "downgrade") {
        parts = append(parts, fmt.Sprintf("%s downgraded %s from %s to %s", 
            rec.Brokerage, rec.Ticker, rec.RatingFrom, rec.RatingTo))
    } else if strings.Contains(strings.ToLower(rec.Action), "initiate") {
        parts = append(parts, fmt.Sprintf("%s initiated coverage on %s with %s rating", 
            rec.Brokerage, rec.Ticker, rec.RatingTo))
    } else {
        parts = append(parts, fmt.Sprintf("%s maintained %s rating on %s", 
            rec.Brokerage, rec.RatingTo, rec.Ticker))
    }
    
    // Target price info
    if rec.TargetTo > 0 && rec.TargetFrom > 0 && rec.TargetTo != rec.TargetFrom {
        change := ((rec.TargetTo - rec.TargetFrom) / rec.TargetFrom) * 100
        if change > 0 {
            parts = append(parts, fmt.Sprintf("Target price raised %.1f%% to $%.2f", change, rec.TargetTo))
        } else {
            parts = append(parts, fmt.Sprintf("Target price lowered %.1f%% to $%.2f", math.Abs(change), rec.TargetTo))
        }
    } else if rec.TargetTo > 0 {
        parts = append(parts, fmt.Sprintf("Price target set at $%.2f", rec.TargetTo))
    }
    
    return strings.Join(parts, ". ")
}

// processRecommendation calcula todos los campos faltantes
func processRecommendation(rec *StockRecommendation) {
    rec.Recommendation = calculateRecommendation(rec)
    rec.Confidence = calculateConfidence(rec)
    rec.Score = calculateScore(rec)
    rec.Reason = generateReason(rec)
}

