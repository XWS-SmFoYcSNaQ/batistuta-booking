import {
  Box,
  Button,
  Checkbox,
  Chip,
  FormControlLabel,
  Slider,
  TextField,
} from "@mui/material";
import { useState } from "react";
import { AccommodationFilter } from "./AccommodationFilter";

const valueLabelFormat = (value: number) => {
  return `${value} EUR`;
};

const calculateValue = (value: number) => {
  return value === 0 ? 0 : 2 ** value;
};

interface ChipData {
  key: number;
  label: string;
}

export const Filters = ({
  setFiltersEnabled,
  setFilters,
}: {
  setFiltersEnabled: (value: boolean) => void;
  setFilters: (filters: AccommodationFilter | null) => void;
}) => {
  const [range, setRange] = useState<number[]>([0, 17]);
  const handleRangeChange = (
    event: Event,
    newValue: number | number[],
    activeThumb: number
  ) => {
    if (!Array.isArray(newValue)) {
      return;
    }
    const minDistance = 2;
    if (newValue[1] - newValue[0] < minDistance) {
      if (activeThumb === 0) {
        const clamped = Math.min(newValue[0], 17 - minDistance);
        setRange([clamped, clamped + minDistance]);
      } else {
        const clamped = Math.max(newValue[1], minDistance);
        setRange([clamped - minDistance, clamped]);
      }
    } else {
      setRange(newValue as number[]);
    }
  };

  const [benefitsData, setBenefitsData] = useState<ChipData[]>([]);

  const handleBenefitDelete = (benefitToDelete: ChipData) => () => {
    setBenefitsData((benefits) =>
      benefits.filter((benefit) => benefit.key !== benefitToDelete.key)
    );
  };

  const [benefitToAdd, setBenefitToAdd] = useState("");
  const addBenefit = (e: any) => {
    e.preventDefault();
    if (benefitToAdd.trim() === "") return;
    const benefit = {
      key:
        (benefitsData.length > 0
          ? benefitsData[benefitsData.length - 1].key
          : -1) + 1,
      label: benefitToAdd,
    };
    benefitsData.push(benefit);
    setBenefitsData(benefitsData);
    setBenefitToAdd("");
  };

  const [distinguished, setDistinguished] = useState(false);

  const submitHandler = (e: any) => {
    e.preventDefault();
    const filter: AccommodationFilter = {
      range: [calculateValue(range[0]), calculateValue(range[1])],
      benefits: benefitsData.map((b) => b.label),
      distinguished,
    };
    setFilters(filter);
  };

  return (
    <Box>
      <Box sx={{ width: 300 }}>
        <Slider
          getAriaLabel={() => "Cost range"}
          value={range}
          onChange={handleRangeChange}
          valueLabelDisplay="auto"
          getAriaValueText={valueLabelFormat}
          valueLabelFormat={valueLabelFormat}
          scale={calculateValue}
          min={0}
          max={17}
          step={1}
        />
      </Box>
      <form onSubmit={addBenefit}>
        <TextField
          required
          label="Benefits"
          value={benefitToAdd ?? ""}
          onChange={(e) => setBenefitToAdd(e.target.value)}
        />
      </form>

      {benefitsData.length > 0 && (
        <Box
          sx={{
            display: "flex",
            flexWrap: "wrap",
            listStyle: "none",
            gap: 1,
            p: 0,
            mt: 2,
          }}
          component="ul"
        >
          {benefitsData.map((data) => (
            <Box key={data.key}>
              <Chip label={data.label} onDelete={handleBenefitDelete(data)} />
            </Box>
          ))}
        </Box>
      )}

      <FormControlLabel
        control={
          <Checkbox
            aria-label="distinguished"
            checked={distinguished}
            onChange={(e) => setDistinguished(e.target.checked)}
          />
        }
        label="Featured Hosts Only"
      />

      <Box sx={{ display: "flex" }}>
        <Button
          type="button"
          color="error"
          onClick={() => setFiltersEnabled(false)}
        >
          Disable Fitlers
        </Button>
        <Button type="button" onClick={submitHandler}>
          Filter
        </Button>
      </Box>
    </Box>
  );
};
